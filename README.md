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

最后更新：2026-03-14 12:21:45 UTC（2026-03-14 20:21:45 UTC+8）

**代理总数：64**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 202.155.12.161:443 | ✓ 1422ms | 否 | ✓ 1007ms | ✓ 1250ms | ✓ 1056ms | http |
| 205.209.118.30:3138 | 否 | ✓ 1959ms | ✓ 1211ms | ✓ 1198ms | ✓ 1167ms | http |
| 162.240.154.26:3128 | ✓ 1045ms | 否 | ✓ 1847ms | 否 | ✓ 1214ms | http |
| 45.167.124.52:8080 | ✓ 1394ms | ✓ 1864ms | 否 | ✓ 1645ms | ✓ 1368ms | http |
| 47.77.193.180:1080 | ✓ 623ms | 否 | ✓ 477ms | ✓ 1526ms | ✓ 725ms | http |
| 59.46.216.131:30001 | ✓ 1843ms | ✓ 1540ms | ✓ 1594ms | 否 | ✓ 1954ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1097ms | ✓ 1286ms | ✓ 1110ms | http |
| 35.225.22.61:80 | 否 | ✓ 1235ms | ✓ 207ms | ✓ 1065ms | 否 | http |
| 45.140.147.155:1081 | ✓ 510ms | ✓ 1862ms | ✓ 1192ms | 否 | 否 | http |
| 45.140.147.155:1082 | ✓ 1119ms | 否 | ✓ 1318ms | 否 | ✓ 1307ms | http |
| 101.43.127.100:8877 | ✓ 1786ms | 否 | ✓ 1895ms | ✓ 1438ms | 否 | http |
| 210.77.29.245:7890 | ✓ 1313ms | ✓ 1568ms | 否 | ✓ 1518ms | 否 | http |
| 150.230.249.50:1080 | ✓ 838ms | 否 | ✓ 910ms | ✓ 1135ms | ✓ 858ms | http |
| 116.80.49.166:3172 | ✓ 1578ms | 否 | 否 | ✓ 1894ms | ✓ 1745ms | http |
| 185.115.74.185:8080 | ✓ 1696ms | ✓ 1648ms | ✓ 1705ms | 否 | 否 | http |
| 81.70.169.194:80 | ✓ 1987ms | 否 | 否 | ✓ 1292ms | ✓ 1487ms | http |
| 45.140.147.82:1081 | ✓ 1031ms | ✓ 1314ms | ✓ 1805ms | ✓ 1583ms | ✓ 1513ms | http |
| 116.80.64.157:7777 | 否 | 否 | ✓ 1575ms | ✓ 1945ms | ✓ 1730ms | http |
| 86.53.183.16:1080 | ✓ 545ms | 否 | ✓ 1164ms | 否 | ✓ 1867ms | http |
| 116.80.65.78:3172 | ✓ 1639ms | 否 | 否 | ✓ 1925ms | ✓ 1853ms | http |
| 165.227.5.10:8888 | ✓ 786ms | ✓ 1515ms | ✓ 1256ms | 否 | 否 | http |
| 62.113.119.14:8080 | 否 | 否 | ✓ 1384ms | ✓ 1794ms | ✓ 1250ms | http |
| 45.136.130.211:8447 | ✓ 919ms | 否 | ✓ 502ms | ✓ 997ms | ✓ 772ms | http |
| 138.124.53.25:7443 | ✓ 1520ms | 否 | ✓ 1635ms | 否 | ✓ 1481ms | http |
| 116.80.65.76:3172 | ✓ 1616ms | 否 | ✓ 1560ms | 否 | ✓ 1862ms | http |
| 85.208.108.43:2094 | ✓ 359ms | 否 | ✓ 558ms | ✓ 1139ms | 否 | http |
| 85.198.96.242:3128 | ✓ 1174ms | 否 | ✓ 1298ms | ✓ 1754ms | ✓ 1379ms | http |
| 45.168.238.193:8443 | 否 | ✓ 1329ms | ✓ 794ms | ✓ 1236ms | ✓ 1004ms | http |
| 45.149.92.147:5001 | ✓ 1261ms | 否 | ✓ 754ms | ✓ 929ms | ✓ 720ms | http |
| 38.210.179.112:999 | ✓ 815ms | 否 | ✓ 1952ms | ✓ 1594ms | 否 | http |
| 101.43.255.96:80 | ✓ 1368ms | 否 | ✓ 1765ms | ✓ 1972ms | ✓ 1360ms | http |
| 216.180.127.45:1080 | ✓ 1107ms | 否 | ✓ 1970ms | ✓ 1919ms | ✓ 1519ms | http |
| 120.92.212.16:8890 | ✓ 1030ms | 否 | 否 | ✓ 1325ms | ✓ 1056ms | http |
| 43.167.227.161:1080 | ✓ 1594ms | 否 | ✓ 559ms | ✓ 1599ms | 否 | http |
| 116.80.49.159:3172 | ✓ 1758ms | 否 | ✓ 1585ms | ✓ 1935ms | 否 | http |
| 101.47.73.135:3128 | ✓ 1394ms | 否 | 否 | ✓ 1987ms | ✓ 1416ms | http |
| 103.39.51.190:8080 | ✓ 1846ms | 否 | 否 | ✓ 1906ms | ✓ 1429ms | http |
| 180.127.149.244:1080 | 否 | ✓ 1288ms | ✓ 1921ms | 否 | ✓ 1003ms | http |
| 45.136.198.40:3128 | ✓ 772ms | ✓ 1987ms | ✓ 1709ms | 否 | ✓ 1647ms | http |
| 157.100.54.4:80 | ✓ 1280ms | 否 | 否 | ✓ 1774ms | ✓ 1483ms | http |
| 114.231.72.199:1080 | ✓ 1145ms | ✓ 1304ms | ✓ 1070ms | 否 | ✓ 1349ms | http |
| 38.145.203.135:8443 | ✓ 640ms | ✓ 1461ms | ✓ 512ms | ✓ 936ms | 否 | http |
| 45.136.131.39:8443 | ✓ 1702ms | ✓ 949ms | ✓ 1062ms | ✓ 1839ms | ✓ 746ms | http |
| 165.225.113.220:10018 | ✓ 1750ms | 否 | ✓ 986ms | ✓ 1350ms | 否 | http |
| 165.225.113.220:11396 | ✓ 1147ms | 否 | ✓ 1174ms | ✓ 1461ms | 否 | http |
| 165.225.113.220:11462 | ✓ 1474ms | 否 | ✓ 1264ms | ✓ 1556ms | ✓ 1426ms | http |
| 165.225.113.220:10958 | 否 | 否 | ✓ 1002ms | ✓ 1486ms | ✓ 1299ms | http |
| 165.225.113.220:11589 | ✓ 1505ms | 否 | ✓ 1044ms | 否 | ✓ 1049ms | http |
| 165.225.113.220:10017 | ✓ 1504ms | 否 | 否 | ✓ 1451ms | ✓ 1019ms | http |
| 165.225.113.220:11191 | ✓ 1468ms | 否 | ✓ 1045ms | ✓ 1454ms | 否 | http |
| 120.92.212.16:7890 | 否 | ✓ 1306ms | 否 | ✓ 1373ms | ✓ 1063ms | http |
| 116.80.49.165:3172 | ✓ 1631ms | 否 | ✓ 1604ms | 否 | ✓ 1737ms | http |
| 165.225.113.220:10462 | ✓ 1503ms | 否 | ✓ 954ms | ✓ 1542ms | ✓ 1254ms | http |
| 165.225.113.220:10829 | ✓ 1509ms | 否 | 否 | ✓ 1448ms | ✓ 997ms | http |
| 165.225.113.220:10525 | ✓ 1531ms | 否 | 否 | ✓ 1449ms | ✓ 997ms | http |
| 165.225.113.220:12004 | ✓ 1519ms | 否 | ✓ 997ms | 否 | ✓ 1090ms | http |
| 165.225.113.220:11584 | ✓ 1527ms | 否 | ✓ 986ms | 否 | ✓ 1100ms | http |
| 165.225.113.220:11918 | ✓ 1516ms | 否 | ✓ 1025ms | ✓ 1494ms | ✓ 1248ms | http |
| 165.225.113.220:11544 | ✓ 1457ms | 否 | ✓ 978ms | 否 | ✓ 1132ms | http |
| 165.225.113.220:11222 | ✓ 1184ms | 否 | 否 | ✓ 1441ms | ✓ 1067ms | http |
| 165.225.113.220:10884 | 否 | 否 | ✓ 1111ms | ✓ 1441ms | ✓ 1107ms | http |
| 165.225.113.220:11143 | ✓ 1173ms | 否 | ✓ 1058ms | 否 | ✓ 1260ms | http |
| 165.225.113.220:11188 | ✓ 1513ms | 否 | ✓ 1044ms | ✓ 1463ms | ✓ 1196ms | http |
| 165.225.113.220:11180 | ✓ 1492ms | 否 | ✓ 1017ms | ✓ 1466ms | ✓ 1193ms | http |

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
