**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-04-12 09:46:06 UTC（2026-04-12 17:46:06 UTC+8）

**代理总数：44**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 44 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 44 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1639ms | 否 | ✓ 905ms | ✓ 1334ms | ✓ 1225ms | http |
| 113.160.132.26:8080 | ✓ 1664ms | 否 | 否 | ✓ 1415ms | ✓ 1169ms | http |
| 212.58.132.5:8888 | ✓ 1780ms | 否 | ✓ 1577ms | 否 | ✓ 1306ms | http |
| 167.103.34.108:8800 | ✓ 1836ms | 否 | 否 | ✓ 1704ms | ✓ 1853ms | http |
| 45.167.124.52:8080 | ✓ 940ms | ✓ 1634ms | ✓ 1307ms | ✓ 1674ms | 否 | http |
| 35.225.22.61:80 | ✓ 314ms | 否 | ✓ 1424ms | 否 | ✓ 874ms | http |
| 43.99.67.68:59409 | ✓ 761ms | ✓ 1118ms | ✓ 813ms | ✓ 969ms | ✓ 789ms | http |
| 167.103.115.102:8800 | ✓ 1219ms | 否 | ✓ 1023ms | ✓ 1162ms | ✓ 1134ms | http |
| 167.103.144.127:8800 | ✓ 1956ms | 否 | ✓ 1667ms | ✓ 1797ms | ✓ 1798ms | http |
| 217.76.245.80:999 | 否 | ✓ 1598ms | ✓ 1227ms | ✓ 1794ms | ✓ 1368ms | http |
| 45.167.125.21:999 | 否 | ✓ 1914ms | ✓ 1256ms | ✓ 1976ms | ✓ 1640ms | http |
| 167.103.31.122:8800 | ✓ 1721ms | 否 | ✓ 1395ms | ✓ 1617ms | ✓ 1468ms | http |
| 185.76.240.254:10001 | ✓ 1118ms | 否 | ✓ 1394ms | 否 | ✓ 1654ms | http |
| 46.39.105.157:8080 | ✓ 1790ms | 否 | ✓ 1797ms | 否 | ✓ 1385ms | http |
| 159.223.225.118:8888 | ✓ 1715ms | 否 | 否 | ✓ 1779ms | ✓ 1890ms | http |
| 210.223.44.230:3128 | ✓ 723ms | ✓ 1387ms | 否 | ✓ 1058ms | 否 | http |
| 147.161.239.240:8800 | ✓ 787ms | ✓ 1692ms | ✓ 1343ms | ✓ 1618ms | ✓ 1556ms | http |
| 115.231.181.40:8128 | ✓ 1017ms | 否 | ✓ 1104ms | ✓ 1640ms | 否 | http |
| 103.113.70.189:1081 | ✓ 321ms | 否 | ✓ 150ms | ✓ 1007ms | ✓ 776ms | http |
| 45.140.147.82:1081 | ✓ 454ms | 否 | ✓ 462ms | ✓ 1431ms | ✓ 1330ms | http |
| 45.140.147.82:1082 | ✓ 1510ms | 否 | ✓ 1126ms | ✓ 1307ms | ✓ 911ms | http |
| 34.96.238.40:8080 | ✓ 1381ms | 否 | ✓ 1320ms | 否 | ✓ 1193ms | http |
| 202.141.161.53:10808 | 否 | ✓ 1605ms | ✓ 1382ms | 否 | ✓ 1109ms | http |
| 93.77.181.116:8888 | ✓ 610ms | 否 | ✓ 929ms | 否 | ✓ 1804ms | http |
| 121.230.9.125:1080 | ✓ 1100ms | ✓ 1452ms | ✓ 1173ms | ✓ 1701ms | ✓ 1231ms | http |
| 129.213.162.27:17777 | ✓ 440ms | 否 | ✓ 1928ms | 否 | ✓ 1555ms | http |
| 107.172.102.234:40621 | ✓ 1264ms | ✓ 993ms | ✓ 1344ms | ✓ 1102ms | ✓ 867ms | http |
| 45.149.92.147:5001 | ✓ 757ms | 否 | ✓ 759ms | ✓ 962ms | ✓ 759ms | http |
| 101.32.244.83:8080 | ✓ 1072ms | 否 | ✓ 1083ms | ✓ 1631ms | ✓ 1438ms | http |
| 121.43.196.210:8222 | ✓ 1140ms | ✓ 1234ms | ✓ 930ms | ✓ 1329ms | ✓ 1050ms | http |
| 121.43.196.213:8222 | ✓ 1052ms | ✓ 1238ms | ✓ 1158ms | ✓ 1268ms | ✓ 1045ms | http |
| 114.55.226.123:10086 | ✓ 1213ms | ✓ 1543ms | ✓ 1133ms | ✓ 1504ms | ✓ 1206ms | http |
| 45.88.0.99:3128 | ✓ 1551ms | 否 | ✓ 1731ms | 否 | ✓ 1840ms | http |
| 103.165.157.206:8088 | ✓ 1978ms | 否 | ✓ 1361ms | ✓ 1577ms | ✓ 1537ms | http |
| 103.134.245.127:8090 | ✓ 1993ms | 否 | ✓ 1784ms | 否 | ✓ 1758ms | http |
| 45.123.142.83:1111 | 否 | 否 | ✓ 1553ms | ✓ 1517ms | ✓ 1438ms | http |
| 121.176.242.215:3084 | ✓ 1028ms | ✓ 1792ms | ✓ 1658ms | 否 | 否 | http |
| 138.124.99.216:8888 | ✓ 676ms | ✓ 1744ms | 否 | 否 | ✓ 1757ms | http |
| 34.231.145.203:7000 | ✓ 88ms | ✓ 1063ms | ✓ 398ms | 否 | ✓ 722ms | http |
| 61.52.131.172:8443 | ✓ 1135ms | ✓ 1261ms | ✓ 1602ms | ✓ 1354ms | ✓ 1092ms | http |
| 38.34.179.173:8452 | ✓ 1852ms | ✓ 1357ms | ✓ 1543ms | 否 | 否 | http |
| 181.78.44.63:999 | ✓ 1538ms | ✓ 1770ms | ✓ 1718ms | ✓ 1912ms | 否 | http |
| 223.16.170.103:80 | ✓ 1277ms | 否 | ✓ 1234ms | ✓ 1228ms | ✓ 1304ms | http |
| 103.39.51.207:8080 | 否 | 否 | ✓ 1924ms | ✓ 1724ms | ✓ 1508ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
