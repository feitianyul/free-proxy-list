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

最后更新：2026-03-04 21:46:42 UTC（2026-03-05 05:46:42 UTC+8）

**代理总数：60**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 59 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 60 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 103ms | ✓ 1080ms | ✓ 1081ms | ✓ 1233ms | ✓ 790ms | http |
| 35.225.22.61:80 | 否 | ✓ 1686ms | 否 | ✓ 1258ms | ✓ 859ms | http |
| 14.56.107.244:3128 | 否 | ✓ 1972ms | ✓ 1167ms | ✓ 1237ms | ✓ 952ms | http |
| 61.72.221.94:3128 | ✓ 1625ms | 否 | 否 | ✓ 1307ms | ✓ 1294ms | http |
| 121.128.121.54:3128 | ✓ 960ms | ✓ 1416ms | ✓ 1289ms | 否 | 否 | http |
| 61.72.221.234:3128 | ✓ 832ms | ✓ 1331ms | 否 | ✓ 1226ms | ✓ 1203ms | http |
| 61.72.221.194:3128 | ✓ 830ms | 否 | ✓ 1452ms | ✓ 1844ms | ✓ 988ms | http |
| 173.212.246.157:3128 | ✓ 1682ms | 否 | ✓ 1531ms | 否 | ✓ 1892ms | http |
| 222.228.171.92:8080 | ✓ 1614ms | 否 | 否 | ✓ 1383ms | ✓ 1170ms | http |
| 138.124.53.25:7443 | ✓ 453ms | 否 | ✓ 1497ms | ✓ 1702ms | ✓ 1419ms | http |
| 61.72.110.94:3128 | 否 | 否 | ✓ 1028ms | ✓ 1880ms | ✓ 958ms | http |
| 61.72.110.54:3128 | ✓ 965ms | ✓ 1669ms | 否 | 否 | ✓ 957ms | http |
| 120.92.212.16:8890 | ✓ 1334ms | ✓ 1477ms | 否 | ✓ 1725ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1285ms | ✓ 1527ms | ✓ 1355ms | ✓ 1756ms | 否 | http |
| 125.128.12.144:3128 | ✓ 1021ms | ✓ 1836ms | ✓ 1269ms | 否 | ✓ 1122ms | http |
| 81.70.169.194:80 | ✓ 1122ms | ✓ 1413ms | ✓ 1127ms | ✓ 1540ms | ✓ 1202ms | http |
| 101.43.255.96:80 | ✓ 1140ms | ✓ 1435ms | ✓ 1151ms | ✓ 1609ms | ✓ 1138ms | http |
| 91.193.240.157:9877 | ✓ 1144ms | 否 | ✓ 1794ms | 否 | ✓ 1697ms | http |
| 125.128.12.14:3128 | 否 | 否 | ✓ 1575ms | ✓ 1953ms | ✓ 965ms | http |
| 35.234.17.221:8080 | 否 | 否 | ✓ 1321ms | ✓ 1352ms | ✓ 1120ms | http |
| 70.22.175.232:3128 | 否 | 否 | ✓ 101ms | ✓ 954ms | ✓ 704ms | http |
| 120.55.163.237:10086 | ✓ 1017ms | ✓ 1293ms | ✓ 1221ms | ✓ 1289ms | ✓ 1035ms | http |
| 46.249.103.192:443 | ✓ 635ms | 否 | ✓ 1513ms | ✓ 1664ms | 否 | http |
| 210.223.44.230:3128 | ✓ 855ms | ✓ 1520ms | 否 | ✓ 1339ms | ✓ 858ms | http |
| 88.80.150.82:8080 | ✓ 1016ms | ✓ 1842ms | 否 | ✓ 1929ms | 否 | https |
| 160.238.65.7:3128 | ✓ 485ms | ✓ 1873ms | ✓ 525ms | 否 | 否 | http |
| 160.238.65.9:3128 | ✓ 756ms | ✓ 1182ms | ✓ 935ms | 否 | 否 | http |
| 160.238.65.6:3128 | ✓ 420ms | ✓ 1257ms | ✓ 1976ms | 否 | 否 | http |
| 160.238.65.2:3128 | ✓ 457ms | ✓ 1465ms | ✓ 1916ms | 否 | 否 | http |
| 160.238.65.8:3128 | ✓ 1778ms | ✓ 1714ms | ✓ 542ms | 否 | 否 | http |
| 160.238.65.3:3128 | ✓ 1452ms | ✓ 1190ms | ✓ 1661ms | ✓ 1625ms | 否 | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 1350ms | ✓ 1114ms | ✓ 756ms | http |
| 15.204.151.144:3128 | ✓ 496ms | ✓ 1716ms | 否 | ✓ 1563ms | 否 | http |
| 14.56.177.44:3128 | ✓ 1002ms | ✓ 1572ms | ✓ 936ms | ✓ 1264ms | ✓ 949ms | http |
| 159.89.31.62:8080 | ✓ 937ms | 否 | ✓ 1662ms | ✓ 1693ms | ✓ 1429ms | http |
| 192.166.82.55:1080 | ✓ 829ms | 否 | ✓ 1332ms | 否 | ✓ 1265ms | http |
| 160.238.65.4:3128 | 否 | 否 | ✓ 1386ms | ✓ 1327ms | ✓ 1761ms | http |
| 160.238.65.5:3128 | 否 | 否 | ✓ 1391ms | ✓ 1990ms | ✓ 1099ms | http |
| 211.171.114.154:3128 | ✓ 1189ms | ✓ 1367ms | ✓ 1609ms | 否 | ✓ 1144ms | http |
| 103.139.138.194:3128 | 否 | 否 | ✓ 1611ms | ✓ 1592ms | ✓ 1513ms | http |
| 103.231.236.202:8182 | ✓ 1970ms | 否 | ✓ 1535ms | ✓ 1919ms | ✓ 1924ms | http |
| 165.225.222.22:11396 | ✓ 386ms | ✓ 1081ms | ✓ 444ms | ✓ 1457ms | ✓ 771ms | http |
| 185.243.218.43:49153 | ✓ 1282ms | 否 | ✓ 1606ms | ✓ 1914ms | ✓ 1587ms | http |
| 45.136.198.40:3128 | ✓ 1037ms | 否 | ✓ 1845ms | 否 | ✓ 1938ms | http |
| 8.137.112.117:3128 | ✓ 1124ms | ✓ 1488ms | ✓ 1203ms | ✓ 1735ms | ✓ 1202ms | http |
| 94.158.49.82:3128 | 否 | ✓ 1879ms | 否 | ✓ 1840ms | ✓ 1616ms | http |
| 45.140.147.155:1081 | ✓ 454ms | ✓ 1916ms | ✓ 815ms | ✓ 1529ms | ✓ 1256ms | http |
| 168.235.110.63:3128 | ✓ 240ms | ✓ 929ms | 否 | ✓ 1023ms | ✓ 1788ms | http |
| 195.123.209.48:3128 | ✓ 1414ms | 否 | ✓ 1740ms | ✓ 1845ms | 否 | http |
| 45.140.147.155:1082 | ✓ 1410ms | 否 | ✓ 1337ms | 否 | ✓ 1219ms | http |
| 24.199.124.151:3128 | ✓ 747ms | ✓ 1647ms | ✓ 1412ms | ✓ 964ms | ✓ 775ms | http |
| 172.212.68.37:3128 | ✓ 701ms | 否 | ✓ 735ms | ✓ 1640ms | ✓ 1060ms | http |
| 5.75.196.26:40000 | ✓ 801ms | 否 | ✓ 622ms | ✓ 1803ms | 否 | http |
| 45.140.147.82:1082 | ✓ 787ms | 否 | ✓ 627ms | 否 | ✓ 933ms | http |
| 45.140.147.82:1081 | ✓ 783ms | ✓ 1389ms | ✓ 1246ms | ✓ 1908ms | 否 | http |
| 90.84.188.97:8000 | ✓ 1864ms | 否 | 否 | ✓ 1620ms | ✓ 1592ms | http |
| 199.38.85.122:40004 | ✓ 966ms | 否 | ✓ 1771ms | 否 | ✓ 1406ms | http |
| 74.48.78.224:2080 | 否 | ✓ 1118ms | ✓ 679ms | ✓ 1463ms | ✓ 937ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1623ms | ✓ 1555ms | ✓ 1578ms | 否 | http |
| 91.233.223.147:3128 | 否 | 否 | ✓ 762ms | ✓ 1888ms | ✓ 1416ms | http |

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
