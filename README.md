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

最后更新：2026-05-02 12:40:30 UTC（2026-05-02 20:40:30 UTC+8）

**代理总数：53**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 53 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 53 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 218.108.131.186:17890 | ✓ 1021ms | ✓ 1294ms | ✓ 1064ms | ✓ 1343ms | ✓ 1071ms | http |
| 113.160.132.26:8080 | ✓ 1928ms | 否 | ✓ 1396ms | ✓ 1531ms | ✓ 1181ms | http |
| 148.230.4.241:999 | ✓ 1074ms | ✓ 1624ms | ✓ 787ms | ✓ 1659ms | ✓ 1469ms | http |
| 45.167.124.71:999 | ✓ 545ms | 否 | ✓ 424ms | ✓ 1584ms | ✓ 1304ms | http |
| 91.184.241.12:443 | ✓ 645ms | 否 | ✓ 1001ms | 否 | ✓ 1887ms | http |
| 72.11.150.178:6005 | ✓ 1187ms | ✓ 1163ms | ✓ 1475ms | 否 | 否 | http |
| 47.85.51.197:1080 | ✓ 685ms | ✓ 875ms | ✓ 227ms | ✓ 903ms | 否 | http |
| 107.150.41.226:18080 | ✓ 500ms | ✓ 1522ms | ✓ 195ms | ✓ 1383ms | ✓ 910ms | http |
| 72.11.151.159:6005 | ✓ 488ms | ✓ 1284ms | 否 | ✓ 1239ms | ✓ 990ms | http |
| 149.51.42.10:8080 | ✓ 1691ms | ✓ 1244ms | 否 | ✓ 1251ms | 否 | http |
| 47.77.216.82:1080 | ✓ 323ms | ✓ 1163ms | 否 | 否 | ✓ 1870ms | http |
| 1.231.81.166:3128 | ✓ 1685ms | ✓ 1222ms | 否 | ✓ 1357ms | ✓ 1614ms | http |
| 120.92.212.16:8890 | ✓ 1257ms | 否 | ✓ 1158ms | ✓ 1894ms | ✓ 1422ms | http |
| 103.157.200.126:3128 | ✓ 1520ms | 否 | ✓ 1076ms | ✓ 1607ms | ✓ 1322ms | http |
| 149.51.42.10:3128 | ✓ 513ms | ✓ 1221ms | 否 | ✓ 1452ms | 否 | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1037ms | ✓ 1472ms | ✓ 1180ms | http |
| 206.206.126.177:2412 | ✓ 1433ms | 否 | ✓ 1555ms | ✓ 1422ms | 否 | http |
| 223.84.151.86:30005 | ✓ 1958ms | ✓ 1909ms | ✓ 1749ms | ✓ 1994ms | ✓ 1600ms | http |
| 43.133.44.89:8888 | ✓ 1858ms | 否 | ✓ 1259ms | 否 | ✓ 964ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1292ms | ✓ 1384ms | ✓ 1391ms | http |
| 86.104.72.219:1081 | ✓ 1553ms | ✓ 1063ms | ✓ 1105ms | 否 | 否 | http |
| 144.31.25.69:21064 | ✓ 1605ms | 否 | ✓ 828ms | 否 | ✓ 1389ms | http |
| 152.32.132.190:7890 | ✓ 1103ms | ✓ 1438ms | ✓ 1956ms | 否 | ✓ 1795ms | http |
| 120.92.212.16:7890 | ✓ 1260ms | 否 | 否 | ✓ 1669ms | ✓ 1354ms | http |
| 152.70.91.193:40000 | ✓ 1864ms | 否 | 否 | ✓ 1751ms | ✓ 1578ms | http |
| 38.180.2.107:3128 | ✓ 830ms | 否 | ✓ 1402ms | 否 | ✓ 1867ms | http |
| 103.247.242.22:8080 | 否 | 否 | ✓ 1516ms | ✓ 1753ms | ✓ 1669ms | http |
| 213.111.146.36:18080 | ✓ 956ms | 否 | ✓ 615ms | 否 | ✓ 1227ms | http |
| 92.119.127.211:6005 | ✓ 932ms | ✓ 1819ms | ✓ 1399ms | ✓ 1857ms | ✓ 1110ms | http |
| 8.154.21.175:3128 | ✓ 1039ms | ✓ 1279ms | ✓ 1039ms | ✓ 1373ms | ✓ 1110ms | http |
| 183.238.3.150:7897 | ✓ 1164ms | ✓ 1395ms | ✓ 1235ms | ✓ 1284ms | ✓ 1036ms | http |
| 121.230.9.96:1080 | ✓ 1473ms | 否 | ✓ 1546ms | 否 | ✓ 1313ms | http |
| 121.230.8.235:1080 | 否 | ✓ 1506ms | 否 | ✓ 1609ms | ✓ 1544ms | http |
| 103.193.145.22:8082 | ✓ 1599ms | 否 | ✓ 1693ms | ✓ 1672ms | ✓ 1610ms | http |
| 121.230.8.55:1080 | ✓ 1352ms | 否 | ✓ 1569ms | 否 | ✓ 1256ms | http |
| 190.12.150.244:999 | ✓ 1283ms | ✓ 1609ms | ✓ 911ms | 否 | 否 | http |
| 152.42.177.32:8888 | ✓ 1532ms | 否 | ✓ 1170ms | ✓ 1536ms | ✓ 1295ms | http |
| 222.107.27.7:8068 | ✓ 1822ms | 否 | 否 | ✓ 1695ms | ✓ 1133ms | http |
| 101.32.243.189:80 | ✓ 1594ms | ✓ 1770ms | ✓ 1866ms | ✓ 1678ms | ✓ 1467ms | http |
| 185.121.13.73:3128 | ✓ 1431ms | 否 | ✓ 985ms | 否 | ✓ 1114ms | http |
| 185.195.71.218:18080 | ✓ 1371ms | ✓ 1584ms | 否 | ✓ 1867ms | ✓ 1554ms | http |
| 120.92.108.86:7890 | 否 | 否 | ✓ 1909ms | ✓ 1929ms | ✓ 1591ms | http |
| 3.101.133.120:80 | ✓ 1488ms | ✓ 1910ms | ✓ 1255ms | ✓ 1472ms | ✓ 996ms | http |
| 62.113.119.14:8080 | ✓ 670ms | 否 | ✓ 886ms | ✓ 1471ms | ✓ 1098ms | http |
| 8.219.97.248:80 | ✓ 1936ms | 否 | ✓ 1736ms | ✓ 1847ms | 否 | http |
| 103.82.23.118:5314 | ✓ 1732ms | 否 | ✓ 1471ms | 否 | ✓ 1406ms | http |
| 34.101.184.164:3128 | ✓ 1838ms | 否 | ✓ 1412ms | ✓ 1456ms | ✓ 1283ms | http |
| 86.104.72.220:1082 | 否 | 否 | ✓ 747ms | ✓ 1996ms | ✓ 721ms | http |
| 86.104.72.220:1081 | 否 | ✓ 996ms | ✓ 1134ms | ✓ 1026ms | ✓ 1047ms | http |
| 45.59.122.132:80 | 否 | ✓ 1925ms | ✓ 984ms | ✓ 1533ms | ✓ 1280ms | http |
| 118.113.247.4:1080 | ✓ 1989ms | 否 | ✓ 1630ms | 否 | ✓ 1531ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1815ms | ✓ 1306ms | ✓ 1699ms | 否 | http |
| 61.52.131.172:8443 | ✓ 1132ms | ✓ 1354ms | ✓ 1113ms | ✓ 1396ms | ✓ 1147ms | http |

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
