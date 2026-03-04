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

最后更新：2026-03-04 03:22:55 UTC（2026-03-04 11:22:55 UTC+8）

**代理总数：55**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 55 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 55 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 166.0.192.117:8888 | ✓ 963ms | ✓ 1125ms | 否 | ✓ 1387ms | ✓ 965ms | http |
| 3.213.157.4:3128 | ✓ 318ms | 否 | ✓ 1391ms | ✓ 1695ms | ✓ 1098ms | http |
| 94.176.3.43:7443 | ✓ 507ms | 否 | ✓ 1303ms | 否 | ✓ 1432ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 754ms | ✓ 926ms | ✓ 1024ms | http |
| 205.209.118.30:3138 | 否 | 否 | ✓ 152ms | ✓ 1182ms | ✓ 899ms | http |
| 190.12.150.244:999 | ✓ 1511ms | 否 | ✓ 1806ms | 否 | ✓ 1584ms | http |
| 59.46.216.131:30001 | ✓ 1024ms | ✓ 1418ms | ✓ 1168ms | ✓ 1428ms | ✓ 1158ms | http |
| 178.156.224.42:3128 | ✓ 1499ms | 否 | ✓ 1768ms | 否 | ✓ 1486ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1394ms | ✓ 1923ms | 否 | ✓ 1312ms | http |
| 5.75.196.26:40000 | ✓ 776ms | ✓ 1357ms | ✓ 1552ms | ✓ 1432ms | ✓ 1788ms | http |
| 90.84.188.97:8000 | ✓ 518ms | 否 | 否 | ✓ 1879ms | ✓ 1315ms | http |
| 120.92.212.16:7890 | ✓ 1072ms | ✓ 1531ms | ✓ 1262ms | ✓ 1306ms | ✓ 1025ms | http |
| 95.85.252.153:21064 | ✓ 862ms | ✓ 1788ms | ✓ 1368ms | 否 | 否 | http |
| 91.193.240.157:9877 | ✓ 1119ms | 否 | ✓ 1225ms | 否 | ✓ 1513ms | http |
| 35.234.17.221:8080 | ✓ 1002ms | ✓ 1873ms | ✓ 1640ms | 否 | 否 | http |
| 103.166.185.54:3128 | ✓ 1793ms | ✓ 1712ms | ✓ 1446ms | ✓ 1329ms | ✓ 1076ms | http |
| 180.103.19.49:1080 | ✓ 1407ms | 否 | ✓ 1220ms | 否 | ✓ 1250ms | http |
| 221.202.27.194:10811 | ✓ 1350ms | 否 | ✓ 1637ms | 否 | ✓ 1539ms | http |
| 4.216.195.194:3128 | ✓ 1156ms | 否 | ✓ 1507ms | ✓ 956ms | ✓ 1022ms | http |
| 111.79.111.126:3128 | ✓ 1607ms | ✓ 1483ms | 否 | ✓ 1943ms | ✓ 1288ms | http |
| 35.225.22.61:80 | ✓ 848ms | ✓ 1132ms | ✓ 468ms | ✓ 1027ms | ✓ 1036ms | http |
| 121.141.161.55:1080 | ✓ 1858ms | ✓ 1271ms | 否 | ✓ 1540ms | 否 | http |
| 103.74.192.243:7890 | 否 | 否 | ✓ 1471ms | ✓ 1139ms | ✓ 1859ms | http |
| 190.9.109.199:999 | ✓ 968ms | ✓ 1332ms | ✓ 1222ms | 否 | 否 | http |
| 190.9.109.194:999 | ✓ 879ms | ✓ 1583ms | ✓ 1171ms | 否 | 否 | http |
| 81.70.169.194:80 | ✓ 1063ms | ✓ 1341ms | ✓ 1102ms | ✓ 1287ms | ✓ 1118ms | http |
| 77.83.203.6:443 | ✓ 1308ms | 否 | ✓ 1544ms | 否 | ✓ 1640ms | http |
| 77.83.203.5:443 | ✓ 1303ms | 否 | ✓ 1684ms | 否 | ✓ 1533ms | http |
| 121.230.9.179:1080 | ✓ 1466ms | 否 | ✓ 1149ms | 否 | ✓ 1904ms | http |
| 94.177.131.12:3128 | 否 | 否 | ✓ 602ms | ✓ 900ms | ✓ 942ms | http |
| 103.215.36.88:18574 | ✓ 1443ms | 否 | ✓ 1214ms | ✓ 1438ms | ✓ 1983ms | http |
| 210.223.44.230:3128 | ✓ 1219ms | 否 | ✓ 1232ms | ✓ 1356ms | ✓ 1732ms | http |
| 106.14.203.63:3333 | 否 | 否 | ✓ 1808ms | ✓ 1229ms | ✓ 937ms | http |
| 121.204.158.249:3128 | ✓ 1483ms | ✓ 1306ms | 否 | ✓ 1404ms | ✓ 1056ms | http |
| 103.215.36.88:15556 | 否 | ✓ 1628ms | ✓ 1279ms | ✓ 1734ms | ✓ 1246ms | http |
| 5.129.237.45:49488 | ✓ 1755ms | ✓ 1570ms | ✓ 1703ms | 否 | 否 | http |
| 159.89.31.62:8080 | ✓ 544ms | 否 | ✓ 1670ms | 否 | ✓ 1396ms | http |
| 103.215.36.88:15247 | ✓ 1647ms | ✓ 1921ms | ✓ 1822ms | 否 | ✓ 1639ms | http |
| 45.129.141.143:3128 | ✓ 1156ms | 否 | ✓ 1638ms | 否 | ✓ 1532ms | http |
| 38.180.2.107:3128 | ✓ 1159ms | ✓ 1986ms | ✓ 1722ms | 否 | ✓ 1881ms | http |
| 103.39.51.190:8080 | ✓ 1894ms | 否 | 否 | ✓ 1686ms | ✓ 1455ms | http |
| 62.113.119.14:8080 | ✓ 1729ms | 否 | ✓ 1432ms | ✓ 1939ms | 否 | http |
| 103.215.36.88:16541 | ✓ 1142ms | 否 | ✓ 1549ms | ✓ 1503ms | ✓ 1478ms | http |
| 74.208.234.198:443 | ✓ 1365ms | ✓ 1153ms | ✓ 1312ms | ✓ 934ms | ✓ 710ms | http |
| 46.249.103.192:443 | ✓ 672ms | 否 | ✓ 1084ms | ✓ 1801ms | 否 | http |
| 192.166.82.55:1080 | ✓ 981ms | ✓ 1502ms | ✓ 547ms | 否 | ✓ 1858ms | http |
| 37.27.100.108:443 | ✓ 1559ms | ✓ 1541ms | ✓ 1715ms | 否 | 否 | http |
| 45.136.198.40:3128 | ✓ 1113ms | 否 | ✓ 1433ms | ✓ 1874ms | ✓ 1755ms | http |
| 125.128.12.14:3128 | 否 | ✓ 1260ms | ✓ 672ms | ✓ 1786ms | ✓ 918ms | http |
| 61.72.221.194:3128 | ✓ 1685ms | ✓ 1188ms | 否 | ✓ 1898ms | ✓ 887ms | http |
| 150.107.140.238:3128 | ✓ 1768ms | 否 | ✓ 955ms | 否 | ✓ 1207ms | http |
| 122.2.48.121:8080 | 否 | 否 | ✓ 1397ms | ✓ 1636ms | ✓ 1411ms | http |
| 101.47.73.135:3128 | ✓ 1926ms | 否 | ✓ 1373ms | 否 | ✓ 988ms | http |
| 162.240.154.26:3128 | ✓ 1489ms | 否 | 否 | ✓ 1472ms | ✓ 1457ms | http |
| 74.242.169.16:3128 | ✓ 1117ms | 否 | ✓ 1199ms | 否 | ✓ 1312ms | http |

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
