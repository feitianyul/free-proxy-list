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

最后更新：2026-04-25 19:50:57 UTC（2026-04-26 03:50:57 UTC+8）

**代理总数：43**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 43 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 43 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 113.160.132.26:8080 | ✓ 1017ms | ✓ 1507ms | ✓ 1029ms | ✓ 1435ms | 否 | http |
| 80.92.204.47:1081 | ✓ 1194ms | ✓ 1291ms | ✓ 1018ms | ✓ 1957ms | ✓ 1530ms | http |
| 177.247.249.5:3128 | ✓ 817ms | ✓ 1379ms | ✓ 1330ms | 否 | 否 | http |
| 206.206.126.177:2412 | ✓ 850ms | 否 | ✓ 1307ms | ✓ 1098ms | 否 | http |
| 47.85.51.197:1080 | ✓ 683ms | ✓ 996ms | ✓ 588ms | ✓ 1387ms | 否 | http |
| 218.108.131.186:17890 | ✓ 892ms | ✓ 1419ms | ✓ 916ms | ✓ 1964ms | ✓ 1451ms | http |
| 120.79.99.232:8099 | 否 | 否 | ✓ 1718ms | ✓ 1796ms | ✓ 1366ms | http |
| 2.27.54.161:1080 | ✓ 823ms | 否 | ✓ 1446ms | ✓ 1771ms | ✓ 1503ms | http |
| 1.231.81.166:3128 | ✓ 1802ms | ✓ 1089ms | ✓ 954ms | ✓ 1126ms | ✓ 828ms | http |
| 35.194.4.51:3128 | ✓ 919ms | ✓ 1258ms | ✓ 976ms | ✓ 1155ms | ✓ 776ms | http |
| 35.225.22.61:80 | ✓ 324ms | 否 | ✓ 1104ms | 否 | ✓ 966ms | http |
| 121.230.8.91:1080 | ✓ 1249ms | ✓ 1400ms | ✓ 1426ms | ✓ 1382ms | ✓ 1071ms | http |
| 121.230.8.72:1080 | ✓ 1107ms | ✓ 1491ms | ✓ 1358ms | ✓ 1822ms | ✓ 1300ms | http |
| 103.155.64.182:8080 | ✓ 1836ms | 否 | ✓ 1541ms | ✓ 1538ms | ✓ 1521ms | http |
| 103.126.238.13:8081 | ✓ 1849ms | 否 | ✓ 1528ms | ✓ 1709ms | ✓ 1737ms | http |
| 36.141.21.200:7890 | ✓ 1018ms | ✓ 1238ms | ✓ 1014ms | ✓ 1331ms | ✓ 1045ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1233ms | ✓ 983ms | ✓ 1282ms | ✓ 1007ms | http |
| 159.89.31.62:8080 | ✓ 523ms | ✓ 1973ms | ✓ 1812ms | ✓ 1877ms | ✓ 1770ms | http |
| 47.112.25.109:7890 | 否 | ✓ 1325ms | 否 | ✓ 1299ms | ✓ 1015ms | http |
| 194.31.87.77:3128 | ✓ 1335ms | 否 | ✓ 1531ms | 否 | ✓ 1668ms | http |
| 120.92.212.16:7890 | ✓ 1013ms | 否 | 否 | ✓ 1503ms | ✓ 1247ms | http |
| 43.133.90.161:8888 | ✓ 1796ms | ✓ 1277ms | 否 | ✓ 1905ms | ✓ 1331ms | http |
| 47.105.98.23:3128 | ✓ 1962ms | 否 | ✓ 1437ms | ✓ 1985ms | 否 | http |
| 62.113.119.14:8080 | ✓ 635ms | ✓ 1726ms | ✓ 649ms | ✓ 1703ms | ✓ 1430ms | http |
| 177.93.132.244:3128 | ✓ 1199ms | 否 | ✓ 792ms | 否 | ✓ 1766ms | http |
| 120.92.108.86:7890 | ✓ 1272ms | 否 | ✓ 1736ms | ✓ 1714ms | ✓ 1384ms | http |
| 183.232.248.73:7890 | ✓ 1385ms | 否 | ✓ 1430ms | ✓ 1467ms | ✓ 1494ms | http |
| 149.62.191.202:3128 | ✓ 1486ms | ✓ 1796ms | ✓ 1519ms | 否 | ✓ 1447ms | http |
| 86.104.74.110:1081 | ✓ 1182ms | ✓ 1604ms | ✓ 1112ms | 否 | ✓ 1556ms | http |
| 8.209.238.110:47701 | ✓ 1442ms | 否 | ✓ 805ms | ✓ 954ms | ✓ 773ms | http |
| 102.36.133.62:8080 | ✓ 1424ms | 否 | ✓ 1492ms | 否 | ✓ 1950ms | http |
| 217.52.247.69:1981 | ✓ 1357ms | ✓ 1920ms | 否 | 否 | ✓ 1982ms | http |
| 47.84.73.61:1080 | ✓ 1606ms | ✓ 1731ms | ✓ 918ms | ✓ 1146ms | ✓ 914ms | http |
| 59.46.216.131:30001 | ✓ 1102ms | ✓ 1442ms | 否 | 否 | ✓ 1096ms | http |
| 45.140.147.82:1081 | ✓ 1571ms | ✓ 1247ms | ✓ 1325ms | 否 | 否 | http |
| 68.183.199.89:1080 | 否 | ✓ 1114ms | 否 | ✓ 1290ms | ✓ 932ms | http |
| 122.53.34.6:8082 | ✓ 1923ms | 否 | 否 | ✓ 1375ms | ✓ 1343ms | http |
| 152.70.91.193:40000 | ✓ 1624ms | 否 | ✓ 1737ms | ✓ 1986ms | ✓ 1464ms | http |
| 61.52.131.172:8443 | ✓ 942ms | ✓ 1296ms | ✓ 987ms | ✓ 1315ms | ✓ 1033ms | http |
| 109.234.38.35:3128 | ✓ 624ms | ✓ 1633ms | ✓ 1739ms | 否 | 否 | http |
| 8.219.195.129:1080 | ✓ 916ms | 否 | ✓ 847ms | ✓ 1150ms | ✓ 912ms | http |
| 103.39.51.207:8080 | ✓ 1500ms | 否 | 否 | ✓ 1559ms | ✓ 1548ms | http |
| 47.84.59.16:1080 | ✓ 1347ms | ✓ 1735ms | ✓ 927ms | 否 | 否 | http |

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
