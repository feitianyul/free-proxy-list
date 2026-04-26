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

最后更新：2026-04-26 00:32:19 UTC（2026-04-26 08:32:19 UTC+8）

**代理总数：47**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 47 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 47 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | ✓ 737ms | ✓ 961ms | ✓ 764ms | ✓ 974ms | 否 | http |
| 218.108.131.186:17890 | ✓ 1020ms | ✓ 1227ms | ✓ 1057ms | ✓ 1597ms | ✓ 1055ms | http |
| 120.92.212.16:8890 | ✓ 1542ms | 否 | ✓ 1337ms | 否 | ✓ 1774ms | http |
| 80.92.204.47:1081 | ✓ 1634ms | ✓ 1299ms | ✓ 1518ms | 否 | ✓ 1626ms | http |
| 2.27.54.161:1080 | ✓ 1221ms | 否 | ✓ 1506ms | 否 | ✓ 1909ms | http |
| 113.160.132.26:8080 | ✓ 1769ms | 否 | ✓ 1998ms | ✓ 1576ms | 否 | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 1454ms | ✓ 1247ms | ✓ 1032ms | http |
| 15.204.233.75:3128 | ✓ 740ms | ✓ 1520ms | ✓ 1832ms | ✓ 1516ms | ✓ 976ms | http |
| 45.76.207.177:40000 | ✓ 890ms | 否 | ✓ 1463ms | ✓ 1546ms | ✓ 1088ms | http |
| 177.93.132.244:3128 | ✓ 1168ms | 否 | ✓ 762ms | 否 | ✓ 1641ms | http |
| 194.87.57.226:3128 | ✓ 679ms | ✓ 1720ms | ✓ 1044ms | 否 | ✓ 1761ms | http |
| 23.95.76.201:8443 | ✓ 900ms | ✓ 1122ms | ✓ 1051ms | ✓ 1245ms | 否 | http |
| 89.35.119.147:3128 | ✓ 681ms | ✓ 1763ms | ✓ 1427ms | ✓ 1896ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1444ms | 否 | ✓ 1600ms | 否 | ✓ 1119ms | http |
| 206.206.126.177:2412 | ✓ 1985ms | 否 | ✓ 1353ms | ✓ 1377ms | ✓ 964ms | http |
| 185.88.101.111:8060 | ✓ 658ms | ✓ 1263ms | ✓ 1532ms | ✓ 1905ms | ✓ 1655ms | http |
| 36.141.21.200:7890 | ✓ 1880ms | 否 | ✓ 1125ms | ✓ 1459ms | ✓ 1134ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1509ms | ✓ 1265ms | ✓ 1528ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1760ms | ✓ 1275ms | ✓ 1932ms | ✓ 1792ms | ✓ 1139ms | http |
| 186.116.148.52:8080 | ✓ 631ms | 否 | ✓ 1796ms | 否 | ✓ 1356ms | http |
| 129.213.162.27:17777 | ✓ 698ms | ✓ 1423ms | ✓ 1983ms | ✓ 1724ms | 否 | http |
| 91.186.217.84:3128 | ✓ 1206ms | 否 | ✓ 1883ms | 否 | ✓ 1613ms | http |
| 103.157.200.126:3128 | ✓ 1953ms | 否 | 否 | ✓ 1581ms | ✓ 1256ms | http |
| 94.158.219.111:3128 | 否 | 否 | ✓ 1315ms | ✓ 1852ms | ✓ 1825ms | http |
| 138.124.81.12:8888 | ✓ 1133ms | 否 | ✓ 603ms | 否 | ✓ 1668ms | http |
| 86.104.74.110:1081 | ✓ 468ms | ✓ 1780ms | ✓ 1320ms | ✓ 1785ms | 否 | http |
| 116.63.160.98:8899 | ✓ 1242ms | ✓ 1449ms | ✓ 1143ms | ✓ 1472ms | ✓ 1299ms | http |
| 8.211.166.184:8081 | 否 | ✓ 1294ms | ✓ 881ms | ✓ 1052ms | ✓ 992ms | http |
| 42.200.76.16:3888 | ✓ 910ms | 否 | ✓ 879ms | ✓ 1072ms | ✓ 876ms | http |
| 159.89.31.62:8080 | ✓ 1050ms | 否 | ✓ 1849ms | 否 | ✓ 1531ms | http |
| 152.42.177.32:8888 | ✓ 1154ms | 否 | ✓ 1807ms | ✓ 1540ms | ✓ 1546ms | http |
| 120.92.108.86:7890 | 否 | 否 | ✓ 1695ms | ✓ 1851ms | ✓ 1845ms | http |
| 117.236.124.166:3128 | ✓ 1288ms | 否 | ✓ 1643ms | 否 | ✓ 1801ms | http |
| 183.232.248.73:7890 | 否 | ✓ 1699ms | ✓ 1272ms | ✓ 1394ms | 否 | http |
| 168.222.254.26:8888 | ✓ 1280ms | 否 | ✓ 1178ms | 否 | ✓ 1345ms | http |
| 213.220.62.63:3128 | ✓ 389ms | ✓ 1050ms | ✓ 683ms | ✓ 1858ms | ✓ 1333ms | http |
| 164.92.148.68:3128 | ✓ 1397ms | ✓ 1910ms | ✓ 615ms | 否 | ✓ 948ms | http |
| 120.79.99.232:8099 | 否 | 否 | ✓ 1525ms | ✓ 1540ms | ✓ 1280ms | http |
| 121.230.8.136:1080 | ✓ 1412ms | ✓ 1605ms | ✓ 1557ms | ✓ 1679ms | ✓ 1273ms | http |
| 45.140.147.155:1081 | ✓ 601ms | ✓ 1327ms | ✓ 1277ms | 否 | ✓ 1090ms | http |
| 84.47.150.125:1080 | ✓ 1939ms | ✓ 1375ms | 否 | 否 | ✓ 1087ms | http |
| 45.88.0.115:3128 | ✓ 396ms | ✓ 1139ms | ✓ 1100ms | ✓ 1954ms | ✓ 1525ms | http |
| 130.61.174.200:1080 | 否 | ✓ 1438ms | ✓ 1958ms | ✓ 1974ms | 否 | http |
| 109.224.242.26:8080 | ✓ 1088ms | ✓ 1814ms | ✓ 1751ms | 否 | 否 | http |
| 103.126.238.13:8081 | 否 | 否 | ✓ 1899ms | ✓ 1838ms | ✓ 1742ms | http |
| 61.52.131.172:8443 | ✓ 1076ms | ✓ 1373ms | ✓ 1080ms | ✓ 1466ms | ✓ 1141ms | http |
| 192.241.132.92:80 | ✓ 801ms | ✓ 1158ms | ✓ 1148ms | ✓ 1220ms | ✓ 977ms | http |

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
