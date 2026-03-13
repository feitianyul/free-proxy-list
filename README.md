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

最后更新：2026-03-13 18:32:38 UTC（2026-03-14 02:32:38 UTC+8）

**代理总数：62**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 64.188.90.36:1080 | ✓ 1301ms | 否 | ✓ 1467ms | 否 | ✓ 1542ms | http |
| 45.167.124.52:8080 | ✓ 772ms | 否 | ✓ 1365ms | ✓ 1935ms | ✓ 1511ms | http |
| 178.236.245.59:3128 | ✓ 824ms | ✓ 1632ms | ✓ 1209ms | ✓ 1523ms | ✓ 1215ms | http |
| 217.76.245.80:999 | ✓ 1062ms | ✓ 1576ms | 否 | ✓ 1385ms | 否 | http |
| 45.136.130.223:8443 | ✓ 640ms | ✓ 1258ms | ✓ 860ms | ✓ 1005ms | ✓ 941ms | http |
| 103.87.67.75:3129 | ✓ 1874ms | 否 | ✓ 1697ms | ✓ 1766ms | ✓ 1469ms | http |
| 119.92.142.80:8082 | ✓ 1578ms | 否 | 否 | ✓ 1617ms | ✓ 1568ms | http |
| 103.82.23.118:5247 | ✓ 1795ms | 否 | ✓ 1538ms | 否 | ✓ 1713ms | http |
| 205.209.118.30:3138 | ✓ 1062ms | 否 | ✓ 969ms | ✓ 1121ms | ✓ 804ms | http |
| 113.160.132.26:8080 | ✓ 1656ms | ✓ 1772ms | 否 | ✓ 1442ms | ✓ 1146ms | http |
| 86.53.183.16:1080 | ✓ 568ms | 否 | ✓ 1505ms | 否 | ✓ 1396ms | http |
| 120.92.212.16:8890 | ✓ 1879ms | 否 | ✓ 1142ms | 否 | ✓ 1140ms | http |
| 81.70.169.194:80 | ✓ 1862ms | ✓ 1445ms | 否 | ✓ 1536ms | ✓ 1171ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 806ms | ✓ 1241ms | ✓ 982ms | http |
| 210.223.44.230:3128 | ✓ 1588ms | ✓ 1740ms | ✓ 1022ms | 否 | ✓ 1131ms | http |
| 101.43.255.96:80 | ✓ 1117ms | ✓ 1619ms | ✓ 1204ms | ✓ 1588ms | ✓ 1940ms | http |
| 14.225.212.37:7890 | ✓ 1761ms | 否 | ✓ 1233ms | ✓ 1464ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1377ms | 否 | ✓ 1384ms | ✓ 1745ms | ✓ 1643ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1612ms | 否 | ✓ 1655ms | ✓ 1316ms | http |
| 106.14.203.63:3333 | 否 | ✓ 1896ms | ✓ 1120ms | 否 | ✓ 1064ms | http |
| 103.86.131.62:80 | ✓ 1394ms | 否 | 否 | ✓ 1613ms | ✓ 1289ms | http |
| 139.162.44.152:3128 | ✓ 896ms | 否 | 否 | ✓ 1244ms | ✓ 1466ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1369ms | ✓ 1542ms | ✓ 1352ms | http |
| 88.80.150.82:8080 | ✓ 1638ms | 否 | ✓ 968ms | 否 | ✓ 1613ms | https |
| 171.251.173.39:5104 | ✓ 1547ms | 否 | ✓ 1850ms | ✓ 1814ms | ✓ 1938ms | http |
| 101.32.244.83:8080 | ✓ 1200ms | 否 | ✓ 1173ms | ✓ 1729ms | ✓ 1565ms | http |
| 121.43.196.213:8222 | ✓ 1122ms | ✓ 1283ms | ✓ 1000ms | ✓ 1397ms | ✓ 1060ms | http |
| 121.43.196.210:8222 | ✓ 1163ms | ✓ 1322ms | ✓ 1051ms | ✓ 1404ms | ✓ 1120ms | http |
| 114.55.226.123:10086 | ✓ 1222ms | 否 | ✓ 1206ms | ✓ 1515ms | ✓ 1202ms | http |
| 40.177.212.89:42741 | ✓ 1087ms | 否 | ✓ 1910ms | ✓ 1942ms | ✓ 1345ms | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 1527ms | ✓ 1535ms | ✓ 1633ms | http |
| 104.248.81.109:3128 | ✓ 876ms | ✓ 1883ms | ✓ 1664ms | ✓ 1520ms | ✓ 1340ms | http |
| 95.3.9.78:3128 | 否 | 否 | ✓ 1761ms | ✓ 1817ms | ✓ 1978ms | http |
| 45.140.147.155:1081 | ✓ 1127ms | ✓ 1500ms | ✓ 1721ms | 否 | ✓ 908ms | http |
| 45.140.147.155:1082 | ✓ 1130ms | ✓ 1575ms | ✓ 1639ms | ✓ 1594ms | ✓ 1328ms | http |
| 47.77.193.180:1080 | ✓ 959ms | 否 | ✓ 424ms | ✓ 945ms | ✓ 696ms | http |
| 207.254.71.62:8088 | ✓ 496ms | ✓ 1661ms | ✓ 1338ms | ✓ 1407ms | ✓ 1465ms | http |
| 159.223.42.219:3128 | ✓ 914ms | 否 | ✓ 1130ms | ✓ 1306ms | ✓ 1041ms | http |
| 165.225.113.220:11589 | 否 | 否 | ✓ 1070ms | ✓ 1262ms | ✓ 996ms | http |
| 165.225.113.220:10155 | ✓ 1588ms | 否 | ✓ 889ms | 否 | ✓ 1002ms | http |
| 165.225.113.220:11584 | ✓ 1587ms | 否 | ✓ 886ms | ✓ 1245ms | 否 | http |
| 165.225.113.220:11143 | ✓ 1587ms | 否 | ✓ 903ms | ✓ 1252ms | 否 | http |
| 165.225.113.220:11845 | ✓ 1588ms | 否 | ✓ 911ms | ✓ 1236ms | ✓ 1004ms | http |
| 45.136.198.40:3128 | ✓ 721ms | 否 | ✓ 1727ms | ✓ 1702ms | ✓ 1433ms | http |
| 165.225.113.220:10884 | 否 | 否 | ✓ 901ms | ✓ 1242ms | ✓ 1011ms | http |
| 165.225.113.220:11180 | ✓ 1587ms | 否 | ✓ 908ms | ✓ 1260ms | 否 | http |
| 147.161.246.38:11814 | ✓ 1417ms | 否 | ✓ 1966ms | ✓ 1967ms | ✓ 1452ms | http |
| 128.199.120.45:9090 | ✓ 1621ms | 否 | ✓ 1478ms | 否 | ✓ 1259ms | http |
| 172.212.68.37:3128 | ✓ 1202ms | 否 | ✓ 715ms | ✓ 1473ms | ✓ 940ms | http |
| 201.150.116.32:999 | ✓ 1798ms | ✓ 1221ms | ✓ 1084ms | ✓ 1421ms | 否 | http |
| 35.225.22.61:80 | ✓ 731ms | 否 | ✓ 1343ms | 否 | ✓ 937ms | http |
| 202.129.206.239:3128 | ✓ 1333ms | 否 | ✓ 1629ms | ✓ 1930ms | ✓ 1890ms | http |
| 128.199.121.61:9090 | ✓ 1813ms | 否 | ✓ 1823ms | 否 | ✓ 1671ms | http |
| 106.14.205.114:483 | ✓ 1250ms | ✓ 1437ms | ✓ 1336ms | ✓ 1331ms | ✓ 1108ms | http |
| 144.31.25.69:21064 | ✓ 1139ms | 否 | ✓ 1016ms | 否 | ✓ 1439ms | http |
| 163.44.126.97:3128 | ✓ 1103ms | 否 | ✓ 1904ms | ✓ 1480ms | ✓ 1179ms | http |
| 111.79.111.126:3128 | ✓ 1581ms | 否 | ✓ 1306ms | 否 | ✓ 1450ms | http |
| 61.52.131.172:8443 | ✓ 1036ms | ✓ 1370ms | ✓ 1070ms | ✓ 1336ms | 否 | http |
| 45.136.131.47:8443 | ✓ 1519ms | ✓ 1383ms | ✓ 252ms | ✓ 906ms | ✓ 800ms | http |
| 103.39.51.190:8080 | ✓ 1706ms | 否 | 否 | ✓ 1665ms | ✓ 1653ms | http |
| 162.240.154.26:3128 | 否 | ✓ 1492ms | ✓ 447ms | ✓ 1109ms | ✓ 1556ms | http |
| 106.117.208.101:7890 | ✓ 1175ms | ✓ 1505ms | 否 | 否 | ✓ 1142ms | http |

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
